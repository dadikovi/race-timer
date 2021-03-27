import React from 'react';
import ActiveGroup from './ActiveGroup';
import CreateSegmentForm from './CreateSegmentForm'
import SegmentCard from './SegmentCard';
import ScannerMock from './ScannerMock';
import { Grid, Paper } from "@material-ui/core";

interface SegmentListState {
  results: any;
  segments: any[];
}

export default class SegmentList extends React.Component {

  state: SegmentListState

  constructor(props: any) {
    super(props);
    this.state = {
      results: {},
      segments: []
    };

    this.getSegments = this.getSegments.bind(this);
    this.refresh = this.refresh.bind(this);
    this.refresh();
  }

  refresh() {
    this.getSegments();
    this.getResults();
  }

  getResults() {
    fetch("http://localhost:8010/race/results") 
    .then(res => res.json())
    .then(
      (result) => {
        this.setState({
          results: result
        })
      }
    )
  }

  getSegments() {
    fetch("http://localhost:8010/segments") 
      .then(res => res.json())
      .then(
        (result) => {
          this.setState({
            segments: result
          })
        }
      )
  }

  render() {
    let segmentCards = []
    let activeGroup = <ActiveGroup></ActiveGroup>

    if (this.state.results && this.state.results.activeGroup) {
      activeGroup = <ActiveGroup participants={this.state.results.activeGroup}></ActiveGroup>
    }

    for (let segment of this.state.segments) {
      segmentCards.push(<SegmentCard segment={segment}></SegmentCard>)
    }

    return(
      <div>
        <Grid container spacing={3}>
          <Grid item xs={6}>
            <Paper style={{padding: '10px'}}>
              <CreateSegmentForm onRefresh={this.refresh} onSegmentCreated={this.getSegments}></CreateSegmentForm>
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
}