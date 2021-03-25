import React, {useState} from 'react';
import { TextField, Button } from "@material-ui/core";

function registerParticipant(startNumber) {
  fetch("http://localhost:8010/participants", {
    method: 'POST',
    body: JSON.stringify({startNumber: parseInt(startNumber)})
  });
}

export default function ScannerMock() {
  const [startingNumber, setStartingNumber] = useState();
  
  return(
    <div>
      <TextField placeholder='Starting number' onChange={e => setStartingNumber(e.target.value)} value={startingNumber}></TextField>
      <Button onClick={() => registerParticipant(startingNumber)} color="primary">Register</Button>
    </div>
  );
}