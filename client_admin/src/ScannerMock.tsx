import React, {useState} from 'react';
import { TextField, Button, Box } from "@material-ui/core";
import { registerParticipant, finishParticipant, useAsyncError } from './service';

interface ScannerMockProps {
  mockChanged: Function
}
export default function ScannerMock(props: ScannerMockProps) {
  const [startingNumber, setStartingNumber] = useState<Number | undefined>();
  const throwError = useAsyncError();
  
  return(
    <Box className="formbox" display="flex" alignItems="center">
      <TextField margin="dense" color = "secondary" variant="outlined" label='Starting number' onChange={e => setStartingNumber(Number(e.target.value))} value={startingNumber}></TextField>
      <Button variant="contained" color="secondary" 
        onClick={ () => {
           registerParticipant(startingNumber)
            .then(() => props.mockChanged())
            .catch((err) => throwError(err));
        } }>
        Register
      </Button>
      <Button variant="contained" color="secondary" 
        onClick={ () => {
          finishParticipant(startingNumber) 
            .then(() => props.mockChanged())
            .catch((err) => throwError(err));
        } }>
        Finish
      </Button>
    </Box>
  );
}