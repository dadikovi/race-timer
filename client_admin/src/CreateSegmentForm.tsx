import { Button, TextField } from '@material-ui/core';
import React, {useState} from 'react';
import { createSegment } from './service';

interface CreateSegmentFormProps {
  onRefresh: Function;
  onSegmentCreated: Function;
}

export default function CreateSegmentForm(props: CreateSegmentFormProps) {

  const [segmentName, setSegmentName] = useState<string | undefined>();

  return(
    <div>
      <TextField variant="filled" label='New segment name' onChange={e => setSegmentName(e.target.value)} value={segmentName}></TextField>
      <Button variant="contained"  onClick={() => createSegment(segmentName, props.onSegmentCreated)} color="primary">Create</Button>
      <Button variant="outlined"  onClick={() => props.onRefresh()} color="primary">Refresh</Button>
    </div>
  );
}