import AdminPanel from './AdminPanel'
import './App.css';
import { AppBar, Toolbar, Typography } from "@material-ui/core";

function App() {
  return (
    <div className="App">
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6">
            race-timer admin client
          </Typography>
        </Toolbar>
      </AppBar>
      <main>
        <AdminPanel></AdminPanel>
      </main>
    </div>
  );
}

export default App;
