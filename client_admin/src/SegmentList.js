import React, { useState } from 'react';
import CreateSegmentForm from './CreateSegmentForm'
import SegmentCard from './SegmentCard';

export default class SegmentList extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      segments: []
    };

    this.getSegments = this.getSegments.bind(this);
    this.getSegments();
  }

  getSegments() {
    fetch("http://localhost:8010/segments") 
      .then(res => res.json())
      .then(
        (result) => {
          this.setState({
            segments: result
          })
        },
        (error) => {
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
        <CreateSegmentForm onSegmentCreated={this.getSegments}></CreateSegmentForm>
        {segmentCards}
      </div>
    );
  }
}