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
  const [autoRefresh, setAutoRefresh] = useState<boolean>(false);
  
  const throwError = useAsyncError();

  const autoRefreshIcon = autoRefresh ? <SyncIcon /> : <SyncDisabledIcon />
  const autoRefreshVariant = autoRefresh ? 'contained' : 'outlined'

  const doAutoRefresh = () => {
    autoRefreshTimer = setTimeout(() => {
      props.onRefresh()
      doAutoRefresh()
    }, 1000);
  }

  const cleanAutoRefresh = () => {
    clearTimeout(autoRefreshTimer)
  }

  const toggleAutoRefresh = () => {
    if (autoRefresh) {
      cleanAutoRefresh()
      setAutoRefresh(false)
    } else {
      setAutoRefresh(true)
      doAutoRefresh()
    }
  }

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

      <ButtonGroup color="primary">
        <Button variant="outlined"  onClick={() => props.onRefresh()}>Refresh</Button>
        <Button variant={autoRefreshVariant} endIcon={autoRefreshIcon} onClick={toggleAutoRefresh}></Button>
      </ButtonGroup>
    </Box>
  );
}