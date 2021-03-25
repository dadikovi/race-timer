import { Button, TextField } from '@material-ui/core';
import React from 'react';

export default class CreateSegmentForm extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      segmentName: null
    };
  }

  createSegment(segmentName) {
    fetch("http://localhost:8010/segments", {
      method: 'POST',
      body: JSON.stringify({name: segmentName})
    }) 
      .then(res => res.json())
      .then(
        () => {
          if (this.props.onSegmentCreated !== undefined) {
            this.props.onSegmentCreated()
          }
        },
        (error) => {
        }
      )
  }

  render() {
    return(
      <div style={{padding: '30px 10px'}}>
        <TextField placeholder='New segment name' onChange={e => this.setState({
          segmentName: e.target.value
        })} value={this.state.segmentName}></TextField>
        <Button onClick={() => this.createSegment(this.state.segmentName)} color="primary">Create</Button>
        <Button onClick={() => this.props.onRefresh()} color="primary">Refresh</Button>
      </div>
    );
  }
}