import { Route, Routes, BrowserRouter } from 'react-router'
import Home from './pages/Home'
import Search from './pages/Search'
import Product from './pages/Product'
import MyAccount from './pages/MyAccount'
// import Payment from './pages/Payment'
import Login from './pages/Login'

const CustomRoutes = () => {
    return (
        <BrowserRouter>
            <Routes>
                <Route path='/' element={<Home/>} />
                <Route path='/search' element={<Search/>} />
                <Route path='/product/:id' element={<Product/>} />
                <Route path='/my-account' element={<MyAccount/>} />
                <Route path='/login' element={<Login/>} />
                {/* <Route path='/payment' element={<Payment/>} /> */}

            </Routes>
        </BrowserRouter>
    )
}

export default CustomRoutes
