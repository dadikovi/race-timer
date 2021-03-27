import { Button, TextField, Box, ButtonGroup, IconButton } from '@material-ui/core';
import React, {useState, useEffect} from 'react';
import SyncDisabledIcon from '@material-ui/icons/SyncDisabled';
import SyncIcon from '@material-ui/icons/Sync';
import { createSegment, useAsyncError } from './service';

interface CreateSegmentFormProps {
  onRefresh: Function;
}

let autoRefreshTimer: NodeJS.Timeout

export default function CreateSegmentForm(props: CreateSegmentFormProps) {

  const [segmentName, setSegmentName] = useState<string | undefined>();  
  const throwError = useAsyncError();

  return(
    <Box className="formbox" display="flex" alignItems="center">
      <TextField margin="dense" variant="outlined" label='New segment name' onChange={e => setSegmentName(e.target.value)} value={segmentName}></TextField>
      <Button variant="contained"  
        onClick={() => {
          createSegment(segmentName)
            .then(() => props.onRefresh())
            .catch((err) => throwError(err));
        }} 
        color="primary">
        Create
      </Button>
    </Box>
  );
}