import './styles.scss';
import Logo from '../../../assets/logo.svg'
import Typography from '@mui/material/Typography';
import Breadcrumbs from '@mui/material/Breadcrumbs';
import Link from '@mui/material/Link';

export default function TopBar() {
  return (
    <div className='footer'>
      <div className='footer-logo'>
        <img src={Logo} alt="logo"/>
      </div>
      <Breadcrumbs aria-label="breadcrumb">
        <Link underline="hover" color="inherit" href="/">
          MUI
        </Link>
        <Link
          underline="hover"
          color="inherit"
          href="/"
        >
          Core
        </Link>
        <Typography sx={{ color:"white" }}>Breadcrumbs</Typography>
      </Breadcrumbs>
    </div>
  );
}
