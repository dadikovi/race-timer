import React, { ErrorInfo } from 'react'
import { Snackbar } from "@material-ui/core";
import Alert from '@material-ui/lab/Alert';

interface ErrorBoundaryState {
    snackOpen: boolean
    hasError: boolean
    error: string
}
export default class ErrorBoundary extends React.Component {
    state: ErrorBoundaryState

    handleClose = () => {
        this.setState({snackOpen: false, error: ''})
    };

    constructor(props: any) {
        super(props);
        this.state = { 
            snackOpen: false,
            hasError: false,
            error: ''
        };
    }
  
    static getDerivedStateFromError(_: Error) {
      return { hasError: false };
    }
  
    componentDidCatch(error: Error, errorInfo: ErrorInfo) {
        console.log(`Yo, I found error: ${error} -- ${errorInfo}`)
        this.setState({hasError: false, snackOpen: true, error: error})
    }
  
    render() {
        return (<div>
            {this.props.children}
            <Snackbar open={this.state.snackOpen} autoHideDuration={6000} onClose={this.handleClose}>
                <Alert onClose={this.handleClose} severity="error">{this.state.error.toString()}</Alert>
            </Snackbar>
            </div>)   
    }
  }