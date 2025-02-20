import './styles.scss'
import { useContext, useEffect, useState } from 'react'

import Layout from '../../components/Layout/Layout'
import TextField from '@mui/material/TextField/TextField'
import Box from '@mui/material/Box';
import { AuthContext } from '../../context/AuthContext';
import { Button } from '@mui/material';
import { BaseUrl } from '../../utils/vars';



// https://picsum.photos/ api for images
const Login = () => {
    const [id, setId] = useState('')
    const [password, setPassword] = useState('')

    const {JWTToken, setJWTToken, customer, setCustomer } = useContext(AuthContext);

    return (
        <Layout>
        <Box
        component="form"
        sx={{ '& .MuiTextField-root': { m: 1, width: '25ch' },
            display: 'flex',
            flexDirection: 'column',
            width: "25%"
        }}
        autoComplete="off"
        >
            <TextField
            required
            id="outlined-required"
            label="Id"
            value={id}
            onChange={e => setId(e.target.value) }
            />
            <TextField
            id="outlined-password-input"
            label="Password"
            type="password"
            autoComplete="current-password"
            onChange={e => setPassword(e.target.value) }
            />

            {/* TODO: this is just a mock login... it logs using the customer id and any password works... */}
            <Button variant="contained"
                onClick={ () => {
                    // fetch product info and put into p
                    const url = new URL('/v1/customer/'+id, BaseUrl).href
                    fetch(url)
                        .then((response) => {
                            return response.json() //FIXME: add error handling
                        }).then((data) => {
                            setCustomer(data);
                            console.log(data)
                            return
                        })
                        
                }}
            >
            Login
            </Button>

        </Box>
        </Layout>
    )
}

export default Login
