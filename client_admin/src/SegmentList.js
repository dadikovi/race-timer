import React from 'react';
import CreateSegmentForm from './CreateSegmentForm'
import SegmentCard from './SegmentCard';

export default class SegmentList extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
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

    for (let segment of this.state.segments) {

      segmentCards.push(<SegmentCard name={segment.name}></SegmentCard>)
    }

    return(
      <div>
        <CreateSegmentForm onRefresh={this.refresh} onSegmentCreated={this.getSegments}></CreateSegmentForm>
        {segmentCards}
      </div>
    );
  }
}