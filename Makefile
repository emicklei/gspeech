# make run flags="-p -i welcome.txt"
run:
	GOOGLE_APPLICATION_CREDENTIALS=service-account-key.json go run *.go $(flags)