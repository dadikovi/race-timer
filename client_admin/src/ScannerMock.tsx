import React, {useState} from 'react';
import { TextField, Button } from "@material-ui/core";

function registerParticipant(startNumber: Number | undefined) {
  if (!startNumber) {
    return
  }
  fetch("http://localhost:8010/participants", {
    method: 'POST',
    body: JSON.stringify({startNumber: startNumber})
  });
}

function finishParticipant(startNumber: Number | undefined) {
  if (!startNumber) {
    return
  }
  fetch(`http://localhost:8010/participants/${startNumber}`, {
    method: 'POST'
  });
}

export default function ScannerMock() {
  const [startingNumber, setStartingNumber] = useState<Number | undefined>();
  
  return(
    <div>
      <TextField variant="filled" label='Starting number' onChange={e => setStartingNumber(Number(e.target.value))} value={startingNumber}></TextField>
      <Button variant="contained" color="secondary" onClick={() => registerParticipant(startingNumber)}>Register</Button>
      <Button variant="contained" color="secondary" onClick={() => finishParticipant(startingNumber)}>Finish</Button>
    </div>
  );
}