import React, { useState } from 'react';
import ActiveGroup from './ActiveGroup';
import CreateSegmentForm from './CreateSegmentForm'
import SegmentCard from './SegmentCard';
import ScannerMock from './ScannerMock';
import { Grid, Paper } from "@material-ui/core";

interface SegmentDto {
  name: string;
  id: Number;
}

interface RaceResultstDo {
  activeGroup: ParticipantDto[];
  segments: SegmentResultsDto[];
}

interface SegmentResultsDto {
  segmentName: string;
  participants: ParticipantDto[];
}

interface ParticipantDto {
  startNumber: Number;
  groupId: Number;
  raceTimeMs: Number;
}

function getResults(setResults: Function) {
  fetch("http://localhost:8010/race/results") 
  .then(res => res.json())
  .then(
    (result) => { setResults(result) }
  )
}

function getSegments(setSegments: Function) {
  fetch("http://localhost:8010/segments") 
    .then(res => res.json())
    .then(
      (result) => { setSegments(result) }
    )
}

function refresh(setResults: Function, setSegments: Function) {
  getResults(setResults);
  getSegments(setSegments);
}

export default function SegmentList(props: any) {

  const [segments, setSegments] = useState<SegmentDto[] | undefined>();
  const [results, setResults] = useState<RaceResultstDo | undefined>();

  // this.refresh();
  
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
              onSegmentCreated={() => getSegments(setSegments)}>
            </CreateSegmentForm>
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