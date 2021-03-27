import React, { useState, useEffect } from 'react';
import ActiveGroup from './ActiveGroup';
import CreateSegmentForm from './CreateSegmentForm'
import SegmentCard from './SegmentCard';
import ScannerMock from './ScannerMock';
import { SegmentDto, RaceResultsDo } from './model';
import { Grid, Paper } from "@material-ui/core";
import {getResults, getSegments} from './service';

function refresh(setResults: Function, setSegments: Function) {
  getResults().then(raceResults => setResults(raceResults));
  getSegments().then(segments => setSegments(segments));  
}

export default function AdminPanel() {

  const [segments, setSegments] = useState<SegmentDto[] | undefined>();
  const [results, setResults] = useState<RaceResultsDo | undefined>();

  const callRefresh = () => refresh(setResults, setSegments);

  useEffect(() => {
    refresh(setResults, setSegments)
  }, []);
  
  let segmentCards = []
  let activeGroup = <ActiveGroup></ActiveGroup>

  if (results && results.activeGroup) {
    activeGroup = <ActiveGroup participants={results.activeGroup}></ActiveGroup>
  }

  if (segments) {
    for (let segment of segments) {
      segmentCards.push(<SegmentCard onRefresh={callRefresh} segment={segment} 
        participants={results?.segments.filter(s => s.segmentName === segment.name)[0].participants}></SegmentCard>)
    }
  }

  return(
    <div>
      <Grid container spacing={3}>
        <Grid item xs={6}>
          <Paper style={{padding: '10px'}}>
            <CreateSegmentForm 
              onRefresh={callRefresh}/>
          </Paper>
        </Grid>
        <Grid item xs={6}>
          <Paper color="error.light" style={{padding: '10px'}}>
            <ScannerMock mockChanged={callRefresh}></ScannerMock>
          </Paper>
        </Grid>
      </Grid>
      
      
      {activeGroup}
      {segmentCards}
    </div>
  );
}