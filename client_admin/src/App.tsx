import AdminPanel from './AdminPanel'
import ErrorBoundary from './ErrorBoundary'
import './App.css';
import { AppBar, Toolbar, Typography, Snackbar } from "@material-ui/core";
import Alert from '@material-ui/lab/Alert';

function App() {
  return (
    <div className="App">
      <ErrorBoundary>
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
      </ErrorBoundary>
    </div>
  );
}

export default App;
