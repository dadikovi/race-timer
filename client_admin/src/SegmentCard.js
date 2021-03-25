import React from 'react';
import { Accordion, AccordionDetails, AccordionSummary, Typography, Button } from "@material-ui/core";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';

function createGroup(segmentId) {
    fetch("http://localhost:8010/groups", {
        method: 'POST',
        body: JSON.stringify({segmentId: segmentId})
    });
}

export default function SegmentCard(props) {
    return(
        <Accordion style={{margin: '5px 0px'}}>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel1a-content"
                id="panel1a-header">
                <Typography>{props.name}</Typography>
            </AccordionSummary>
            <AccordionDetails>
                <Button onClick={() => createGroup(props.id)} color="primary">Start new group</Button>
            </AccordionDetails>
        </Accordion>
    );
}