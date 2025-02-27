# API Exercise de boot.dev

Learn HTTP Servers in Go

https://www.boot.dev/courses/learn-http-servers-golang


Une API Serveur qui permet la gestion sécurisée de Chirps (des tweets fictifs) avec la gestion d'utilisateurs, clés JWS et autres mots de passes


	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	mux.HandleFunc("POST /api/chirps", apiCfg.handlerChirps)
	mux.HandleFunc("GET /api/chirps", apiCfg.handlerAllChirps)
	mux.HandleFunc("GET /api/chirps/{id}", apiCfg.handlerGetOneChirp)
	mux.HandleFunc("DELETE /api/chirps/{id}", apiCfg.handlerDeleteChirp)

	mux.HandleFunc("POST /api/users", apiCfg.handlerUsers)
	mux.HandleFunc("PUT /api/users", apiCfg.handlerEditUsers)

	mux.HandleFunc("POST /api/login", apiCfg.handlerUsersLogin)
	mux.HandleFunc("POST /api/refresh", apiCfg.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", apiCfg.handlerRevoke)

	mux.HandleFunc("POST /api/polka/webhooks", apiCfg.handlerPolkaHooks)