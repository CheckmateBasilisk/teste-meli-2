import './styles.scss'
import Paper from '@mui/material/Paper';

import Divider from '@mui/material/Divider';

import Layout from '../../components/Layout/Layout'
import AccountActions from './AccountActions/AccountActions'
import AccountSummary from './AccountSummary/AccountSummary'

import { useContext, useEffect } from 'react';
import { AuthContext } from '../../context/AuthContext';

//const MyAccount = (p : param) => {
const MyAccount = () => {
    const {JWTToken, setJWTToken, customer, setCustomer } = useContext(AuthContext);

    return (
        <Layout>

            {JSON.stringify(customer)}

            <Paper elevation={3}>
                <AccountSummary CustomerEmail={customer.email} CustomerLogin={customer.login} CustomerName={customer.name}/>
                <Divider variant="middle" flexItem />
                <AccountActions />
            </Paper>
        </Layout>
    )
}

export default MyAccount
