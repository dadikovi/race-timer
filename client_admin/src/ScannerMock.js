import React, {useState} from 'react';
import { TextField, Button } from "@material-ui/core";

function registerParticipant(startNumber) {
  fetch("http://localhost:8010/participants", {
    method: 'POST',
    body: JSON.stringify({startNumber: parseInt(startNumber)})
  });
}

function finishParticipant(startNumber) {
  fetch(`http://localhost:8010/participants/${parseInt(startNumber)}`, {
    method: 'POST'
  });
}

export default function ScannerMock() {
  const [startingNumber, setStartingNumber] = useState();
  
  return(
    <div>
      <TextField variant="filled" label='Starting number' onChange={e => setStartingNumber(e.target.value)} value={startingNumber}></TextField>
      <Button variant="contained" color="secondary" onClick={() => registerParticipant(startingNumber)}>Register</Button>
      <Button variant="contained" color="secondary" onClick={() => finishParticipant(startingNumber)}>Finish</Button>
    </div>
  );
}