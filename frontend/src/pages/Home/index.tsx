import './styles.scss'
import { useState } from 'react'

import Layout from '../../components/Layout/Layout'
import landingDog from '../../assets/landing-dog.webp'


// https://picsum.photos/ api for images
const Home = () => {
    //const[state, setState] = useState(0)

    return (
        <Layout>
            <div style={{display: 'flex',justifyContent: 'center',alignItems: 'center',}}>
                <img src={landingDog} alt={'cute dog at my landing page!'} />
            </div>
        </Layout>
    )
}

export default Home
