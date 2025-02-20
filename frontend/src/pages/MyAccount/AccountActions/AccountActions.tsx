import './styles.scss'
import { useState } from 'react'
import Grid2 from '@mui/material/Grid2'
import Paper from '@mui/material/Paper';

import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import Divider from '@mui/material/Divider';

import AutoStoriesIcon from '@mui/icons-material/AutoStories';
import GiteIcon from '@mui/icons-material/Gite';
import EditIcon from '@mui/icons-material/Edit';

// https://picsum.photos/ api for images
const AccountActions = () => {
    const[state, setState] = useState(0)

    return (
        <List>
        <ListItemButton>
            <ListItemIcon>
                <EditIcon />
            </ListItemIcon>
            <ListItemText primary="Edit Info" />
        </ListItemButton>
        <ListItemButton>
            <ListItemIcon>
                <AutoStoriesIcon />
            </ListItemIcon>
            <ListItemText primary="Purchase History" />
        </ListItemButton>
        <ListItemButton>
            <ListItemIcon>
                <GiteIcon />
            </ListItemIcon>
            <ListItemText primary="Addresses" />
        </ListItemButton>
        </List>
    )
}

export default AccountActions
