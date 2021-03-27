import { Button, TextField } from '@material-ui/core';
import React, {useState} from 'react';
import { createSegment, useAsyncError } from './service';

interface CreateSegmentFormProps {
  onRefresh: Function;
}

export default function CreateSegmentForm(props: CreateSegmentFormProps) {

  const [segmentName, setSegmentName] = useState<string | undefined>();
  const throwError = useAsyncError();

  return(
    <div>
      <TextField variant="filled" label='New segment name' onChange={e => setSegmentName(e.target.value)} value={segmentName}></TextField>
      <Button variant="contained"  
        onClick={() => {
          createSegment(segmentName)
            .then(() => props.onRefresh())
            .catch((err) => throwError(err));
        }} 
        color="primary">
        Create
      </Button>
      <Button variant="outlined"  onClick={() => props.onRefresh()} color="primary">Refresh</Button>
    </div>
  );
}