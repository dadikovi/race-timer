import React, { useState, useEffect } from 'react';
import ActiveGroup from './ActiveGroup';
import CreateSegmentForm from './CreateSegmentForm'
import SegmentCard from './SegmentCard';
import ScannerMock from './ScannerMock';
import RefreshForm from './RefreshForm';
import { SegmentDto, RaceResultsDo } from './model';
import { Box, CardContent, Divider, Card } from "@material-ui/core";
import {getResults, getSegments, useAsyncError} from './service';

function refresh(setResults: Function, setSegments: Function): Promise<any> {
  getSegments().then(segments => setSegments(segments));  
  return getResults()
    .then(raceResults => setResults(raceResults))
}

interface AdminPanelProps {
  displayAdminFeatures: boolean
}

export default function AdminPanel(props: AdminPanelProps) {

  const [segments, setSegments] = useState<SegmentDto[] | undefined>();
  const [results, setResults] = useState<RaceResultsDo | undefined>();
  const throwError = useAsyncError();

  const callRefresh = () => refresh(setResults, setSegments).catch(err => throwError(err))

  useEffect(() => {
    refresh(setResults, setSegments)
  }, []);
  
  let segmentCards = []
  if (segments) {
    for (let segment of segments) {
      segmentCards.push(<SegmentCard key={segment.name} onRefresh={callRefresh} segment={segment} 
        participants={results?.segments?.filter(s => s.segmentName === segment.name)[0]?.participants}></SegmentCard>)
    }
  }

  return(
    <div>
      {props.displayAdminFeatures && <Card elevation = {3}>
        <CardContent>
          <Box className="formbox" display="flex" alignItems="center">
            <CreateSegmentForm 
              onRefresh={callRefresh}/>
            <Divider orientation="vertical" flexItem />
            <RefreshForm 
              onRefresh={callRefresh}/>
            <Divider orientation="vertical" flexItem />
            <ScannerMock mockChanged={callRefresh}></ScannerMock>
          </Box>
        </CardContent>
        </Card>}
      
      {props.displayAdminFeatures && <ActiveGroup onRefresh={callRefresh} data={results?.activeGroup}></ActiveGroup>}
      
      {segmentCards}
    </div>
  );
}