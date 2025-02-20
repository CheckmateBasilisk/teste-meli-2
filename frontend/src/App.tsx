import { CartContextProvider } from './context/CartContext.tsx'
import { AuthContextProvider } from './context/AuthContext.tsx'
import CustomRoutes from './routes.tsx'

function App() {

  return (
    <AuthContextProvider>
    <CartContextProvider>
      <CustomRoutes/>
    </CartContextProvider>
    </AuthContextProvider>
  )
}

export default App
