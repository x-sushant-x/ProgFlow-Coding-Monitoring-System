import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import { Auth0Provider } from '@auth0/auth0-react'
import { MyContextProvider } from './contexts/Context.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <MyContextProvider>
    <Auth0Provider
    // Add Auth0 Details Here
      domain=""
      clientId=""
      authorizationParams={{
        redirect_uri: window.location.origin
      }}
    >
      <App />
    </Auth0Provider>
  </MyContextProvider>
)
