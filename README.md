## ProgFlow - Coding Monitoring System

ProgFlow is an open-source self-hosted software to keep track of your coding time directly from VS Code. This source code consist of backend server, frontend web application and database schema.

### Written In
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white) ![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)

<!-- ### How It Works?
1. I have developed an extension that will be integrated into your VS Code.
2. Now you have to create an account on web portal -->

### How To Setup?
To setup ProgFlow we have to get 4 things running. These four things are Database, Golang Server, Web Frontend and VS Code Extension. Let's setup them one by one.

### Database Schema
Find `db-schema.sql` file under `db` folder and import it into your PostgreSQL Database.

### Server
1. Install <a href = 'https://go.dev/'> Golang</a> in your computer.
2. Go to `server/db/connect.go` and change your database details there.
3. <b>Important</b> you can use environment variables to store this sensitive data. To make things simple I haven't. But it's highly recommended.
4. Run `make run` command inside `/server` folder.
5. Your backend server is now setup and running.

### Frontend
1. Go inside `web/frontend/src/main.tsx` and add Auth0 Domain and ClientID there. 
2. Simply run command `npm run dev`.
3. Now go to `localhost:5173` and Sign Up.
4. Your will then be redirected to Home Page. Now go to postgres database and copy your `API Key` from `users` table.


### Extension
1. It's source code is present in `progflow-vscode-extension` and can build a installer using <a href = 'https://code.visualstudio.com/api/working-with-extensions/publishing-extension'> this </a> documentation.
2. You will get vsix that can be installed directly in VS Code.
3. After installing this extension press `CTRL + SHIFT + P` and search for `Set API Key` and paste your API Key there.
4. Your extension is now connected with server and you can see your coding activity on `localhost:5173`

### Final Words
If you face any issue while setting up ProgFlow please open an issue. I'll reply as soon as possible.