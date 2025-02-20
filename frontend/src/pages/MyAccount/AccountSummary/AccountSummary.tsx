import './styles.scss'

import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';

import ClientAvatar from '../../../components/ClientAvatar/ClientAvatar';


interface param {
    //children : React.ReactNode
    CustomerName: string
    CustomerLogin: string
    CustomerEmail: string
}


// https://picsum.photos/ api for images
const AccountSummary = (p : param) => {
    return (
        <List>
            <ListItem>
                <ListItemIcon>
                    <ClientAvatar/>
                </ListItemIcon>
                <ListItemText primary={p.CustomerName} />
            </ListItem>
            <ListItem>
                <ListItemText primary={p.CustomerLogin} />
            </ListItem>
            <ListItem>
                <ListItemText primary={p.CustomerEmail} />
            </ListItem>
        </List>
    )
}

export default AccountSummary
