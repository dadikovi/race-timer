import { Button, TextField } from '@material-ui/core';
import React, { useState } from 'react';
import CreateSegmentForm from './CreateSegmentForm'

export default class SegmentList extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      segments: []
    };

    this.getSegments = this.getSegments.bind(this);
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
    return(
      <div>
        <CreateSegmentForm onSegmentCreated={this.getSegments}></CreateSegmentForm>
        {JSON.stringify(this.state.segments)}
      </div>
    );
  }
}