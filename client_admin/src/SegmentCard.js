import React from 'react';
import { Accordion, AccordionSummary, Typography } from "@material-ui/core";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

export default class SegmentCard extends React.Component {
    
    render() {
        return(
          <div>
            <Accordion>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel1a-content"
                id="panel1a-header">
                <Typography>{this.props.name}</Typography>
                </AccordionSummary>
            </Accordion>
          </div>
        );
      }
}