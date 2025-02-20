import './style.scss'
import TopBar from './TopBar/TopBar'
import Footer from './Footer/Footer'
import Container from '@mui/material/Container'


interface param {
    children : React.ReactNode
}

const Layout = (p : param) => {
    return (
        <main>
            <TopBar/>
            <Container maxWidth="lg" sx={{padding: "0.5rem"}}>
                {p.children}
            </Container>
            <Footer/>
        </main>
    )
}

export default Layout