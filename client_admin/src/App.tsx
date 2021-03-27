import AdminPanel from './AdminPanel'
import ErrorBoundary from './ErrorBoundary'
import './App.css';
import { AppBar, Toolbar, Typography, IconButton } from "@material-ui/core";
import PowerOffIcon from '@material-ui/icons/PowerOff';
import PowerIcon from '@material-ui/icons/Power';
import { useState } from 'react';

function App() {

  let [adminOn, switchAdmin] = useState<boolean>(true)

  const switchAdminIcon = adminOn ? <PowerIcon /> : <PowerOffIcon />
  const onAdminSwitched = () => { 
    if (adminOn) {
      switchAdmin(false)
    } else {
      switchAdmin(true)
    }
  }


  return (
    <div className="App">
      <ErrorBoundary>
        <AppBar position="static">
        <Toolbar>
        <IconButton onClick={onAdminSwitched} edge="start" color="inherit" aria-label="menu">
          {switchAdminIcon}
        </IconButton>
          {adminOn &&<Typography variant="h6">
            race-timer admin client
          </Typography>}
        </Toolbar>
      </AppBar>
      <main>
        <AdminPanel displayAdminFeatures={adminOn}></AdminPanel>
      </main>
      </ErrorBoundary>
    </div>
  );
}

export default App;
