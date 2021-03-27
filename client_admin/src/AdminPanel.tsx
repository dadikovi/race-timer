import React, { useState, useEffect } from 'react';
import ActiveGroup from './ActiveGroup';
import CreateSegmentForm from './CreateSegmentForm'
import SegmentCard from './SegmentCard';
import ScannerMock from './ScannerMock';
import RefreshForm from './RefreshForm';
import { SegmentDto, RaceResultsDo } from './model';
import { Grid, Paper } from "@material-ui/core";
import {getResults, getSegments} from './service';

function refresh(setResults: Function, setSegments: Function) {
  getResults().then(raceResults => setResults(raceResults));
  getSegments().then(segments => setSegments(segments));  
}

interface AdminPanelProps {
  displayAdminFeatures: boolean
}

export default function AdminPanel(props: AdminPanelProps) {

  const [segments, setSegments] = useState<SegmentDto[] | undefined>();
  const [results, setResults] = useState<RaceResultsDo | undefined>();

  const callRefresh = () => refresh(setResults, setSegments);

  useEffect(() => {
    refresh(setResults, setSegments)
  }, []);
  
  let segmentCards = []
  if (segments) {
    for (let segment of segments) {
      segmentCards.push(<SegmentCard key={segment.name} onRefresh={callRefresh} segment={segment} 
        participants={results?.segments.filter(s => s.segmentName === segment.name)[0].participants}></SegmentCard>)
    }
  }

  return(
    <div>
      {props.displayAdminFeatures && <Grid container spacing={1}>
        <Grid item xs={4}>
          <Paper style={{padding: '10px'}}>
            <CreateSegmentForm 
              onRefresh={callRefresh}/>
          </Paper>
        </Grid>
        <Grid item xs={4}>
          <Paper style={{padding: '10px'}}>
            <RefreshForm 
              onRefresh={callRefresh}/>
          </Paper>
        </Grid>
        <Grid item xs={4}>
          <Paper color="error.light" style={{padding: '10px'}}>
            <ScannerMock mockChanged={callRefresh}></ScannerMock>
          </Paper>
        </Grid>
      </Grid>}
      
      {props.displayAdminFeatures && <ActiveGroup participants={results?.activeGroup}></ActiveGroup>}
      
      {segmentCards}
    </div>
  );
}