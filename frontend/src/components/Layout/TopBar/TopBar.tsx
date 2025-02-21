import InputBase  from '@mui/material/Input';
import './styles.scss';
import SearchIcon from '@mui/icons-material/Search';
import Logo from '../../../assets/logo.svg'
import ClientAvatar from '../../ClientAvatar/ClientAvatar';
import ShoppingCart from '../ShoppingCart/ShoppingCart';
import IconButton from '@mui/material/IconButton';
import { useNavigate } from 'react-router';
import { useState } from 'react';

export default function TopBar() {
  const navigate = useNavigate()
  const [searchContent, setSearchContent] = useState('')
  return (

    <div className='top-bar'>
      <div className='top-bar-logo'>
        <IconButton aria-label='My Account' onClick={()=>{navigate("/")}}>
          <img src={Logo} alt="logo"/>
        </IconButton>

      </div>
      <div className='top-bar-search'>
        <div className='top-bar-search-container'>
          <InputBase className='top-bar-search-input' placeholder='Search Products'onChange={(e) => {setSearchContent(e.target.value)}} />

          {/* FIXME: this doesnt change the url... */}
          <IconButton className='top-bar-search-icon' onClick={()=>{navigate("/search?name="+searchContent)}}>
              <SearchIcon />
          </IconButton>

        </div>
      </div>
      <div className='top-bar-right'>

        <IconButton aria-label='My Account' onClick={()=>{navigate("/my-account")}}>
          <ClientAvatar/>
        </IconButton>
        
        <ShoppingCart/>
      </div>
    </div>
  );
}
