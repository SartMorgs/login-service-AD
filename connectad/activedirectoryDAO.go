package connectad

import(
	"fmt"
	"strings"
	"github.com/go-ldap/ldap/v3"
)

type ConnectAD struct{
	loginUsername string
	loginPassword string
}

var(
	UserCredentials ConnectAD
	try_count 		int 	= 0
	flag_block		bool 	= false
)


const(
	ldapServer = "your_server:port_number"
	// Filter all users without desatived accounts
	filterDN = "(&(objectCategory=user)(|(UserAccountControl=512)(UserAccountControl=544)(UserAccountControl=66048))(&(sAMAccountName={username})))"
	// For a example.com domain
	baseDN = "DC=example,DC=com"
)

func SetConnectionAD(username, password string){
	flag_block = false
	try_count = 0

	UserCredentials = ConnectAD{
		loginUsername: username,
		loginPassword: password,
	}
}

func AllowTry(){
	try_count = 0
	flag_block = false
}

func Connect() (*ldap.Conn, error){

	if flag_block != false{
		return nil, fmt.Errorf("Try Enough! 30s Blocked")
	}

	ldapBind := UserCredentials.loginUsername + "@" + ldapServer

	conn, err := ldap.Dial("tcp", ldapServer)

	if err != nil{
		return nil, fmt.Errorf("Failed to connect. %s", err)
	}

	if err := conn.Bind(ldapBind, UserCredentials.loginPassword); err != nil{
		if try_count < 3{
			try_count += 1
			return nil, fmt.Errorf("Failed to bind. %s, err")
		}else{
			return nil, fmt.Errorf("Try Enough! 30s Blocked")
		}
	}

	return conn, nil
}

func Auth(conn *ldap.Conn) error{
	result, err := conn.Search(ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter(UserCredentials.loginUsername),
		[]string{"dn"},
		nil,
	))

	if err != nil{
		return fmt.Errorf("Failed to find user. %s", err)
	}

	if len(result.Entries) < 1{
		return fmt.Errorf("User does not exist")
	}

	if len(result.Entries) > 1{
		return fmt.Errorf("Too many entries returned")
	}

	if err := conn.Bind(result.Entries[0].DN, UserCredentials.loginPassword); err != nil{
		fmt.Printf("Failed to auth. %s", err)
	} else{
		fmt.Printf("Authenticated successfully!")
	}

	return nil
}


func filter(needle string) string {
	res := strings.Replace(
		filterDN,
		"{username}",
		needle,
		-1,
	)

	return res
}