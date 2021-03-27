import { Button, TextField, Box, ButtonGroup, IconButton } from '@material-ui/core';
import React, {useState, useEffect} from 'react';
import SyncDisabledIcon from '@material-ui/icons/SyncDisabled';
import SyncIcon from '@material-ui/icons/Sync';
import { createSegment, useAsyncError } from './service';

interface RefreshFormProps {
  onRefresh: Function;
}

let autoRefreshTimer: NodeJS.Timeout

export default function RefreshForm(props: RefreshFormProps) {
  const [autoRefresh, setAutoRefresh] = useState<boolean>(false);

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
    <ButtonGroup color="primary">
        <Button variant="outlined"  onClick={() => props.onRefresh()}>Refresh</Button>
        <Button variant={autoRefreshVariant} endIcon={autoRefreshIcon} onClick={toggleAutoRefresh}></Button>
    </ButtonGroup>
  );
}