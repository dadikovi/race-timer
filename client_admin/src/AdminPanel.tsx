import React, { useState, useEffect } from 'react';
import ActiveGroup from './ActiveGroup';
import CreateSegmentForm from './CreateSegmentForm'
import SegmentCard from './SegmentCard';
import ScannerMock from './ScannerMock';
import { SegmentDto, RaceResultsDo } from './model';
import { Grid, Paper } from "@material-ui/core";
import {getResults, getSegments} from './service';

async function refresh(setResults: Function, setSegments: Function) {
  const raceResults = await getResults();
  setResults(raceResults)
  const segments = await getSegments();
  setSegments(segments)
}

export default function AdminPanel() {

  const [segments, setSegments] = useState<SegmentDto[] | undefined>();
  const [results, setResults] = useState<RaceResultsDo | undefined>();

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
      segmentCards.push(<SegmentCard segment={segment}></SegmentCard>)
    }
  }

  return(
    <div>
      <Grid container spacing={3}>
        <Grid item xs={6}>
          <Paper style={{padding: '10px'}}>
            <CreateSegmentForm 
              onRefresh={() => refresh(setResults, setSegments)} 
              onSegmentCreated={async () => {
                const segments = await getSegments();
                setSegments(segments)
              }} />
          </Paper>
        </Grid>
        <Grid item xs={6}>
          <Paper style={{padding: '10px'}}>
            <ScannerMock></ScannerMock>
          </Paper>
        </Grid>
      </Grid>
      
      
      {activeGroup}
      {segmentCards}
    </div>
  );
}