import { Button, TextField } from '@material-ui/core';
import React, {useState} from 'react';

function createSegment(segmentName: string | undefined, onSegmentCreated: Function | undefined) {
  fetch("http://localhost:8010/segments", {
    method: 'POST',
    body: JSON.stringify({name: segmentName})
  }) 
    .then(res => res.json())
    .then(
      () => {
        if (onSegmentCreated !== undefined) {
          onSegmentCreated()
        }
      },
      (error) => {
      }
    )
}

export default function CreateSegmentForm(props: any) {

  const [segmentName, setSegmentName] = useState<string | undefined>();

  return(
    <div>
      <TextField variant="filled" label='New segment name' onChange={e => setSegmentName(e.target.value)} value={segmentName}></TextField>
      <Button variant="contained"  onClick={() => createSegment(segmentName, props.onSegmentCreated)} color="primary">Create</Button>
      <Button variant="outlined"  onClick={() => props.onRefresh()} color="primary">Refresh</Button>
    </div>
  );
}