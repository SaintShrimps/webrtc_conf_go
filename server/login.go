package server

/*
func LoginReguestHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Print("HELLO FELLOW KIDS ", r.Body)
	decoder := json.NewDecoder(r.Body)
	var user js_LP
	err := decoder.Decode(&user)
	if err != nil {
		log.Print("JSON faild: ", err)
	}

	log.Print(user)


	connPG, err := pgx.Connect(context.Background(), URLDatabase)

	if err != nil {
		log.Print("Unable to connect to database:", err)
	}
	defer connPG.Close(context.Background())

	var id_user uint16
	err = connPG.QueryRow(
		context.Background(), 
		"select id from public.users where login=$1 and pass=$2 ", user.login, user.password).Scan(id_user)
	if err != nil {
		log.Print("QueryRow failed: ", err)
	}else{
		log.Print("gooood job")
	}

	// login, ok := r.URL.Query()["login"]
	// if !ok {
	// 	log.Print("roomID missing in URL Parameters")
	// 	return
	// }

	// pass, ok := r.URL.Query()["pass"]
	// if !ok {
	// 	log.Print("roomID missing in URL Parameters")
	// 	return
	// }

	// log.Print("login: %s Pass: %s", login, pass)
	//
	//
	//SELECT login, "passWord", id FROM public.users;
	// type resp struct {
	// 	access string `json:"access"`
	// } 

}
*/