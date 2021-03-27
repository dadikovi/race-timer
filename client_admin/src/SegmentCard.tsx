import React from 'react';
import { Accordion, AccordionDetails, AccordionSummary, Typography, Button } from "@material-ui/core";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import { SegmentDto } from './model';

function createGroup(segmentId: Number | undefined) {
    fetch("http://localhost:8010/groups", {
        method: 'POST',
        body: JSON.stringify({segmentId: segmentId})
    });
}

interface SegmentCardProps {
    segment: SegmentDto;
}
export default function SegmentCard(props: SegmentCardProps) {
    console.log(JSON.stringify(props))
    return(
        <Accordion style={{margin: '5px 0px'}}>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel1a-content"
                id="panel1a-header">
                <Typography>{props.segment.name}</Typography>
            </AccordionSummary>
            <AccordionDetails>
                <Button variant="contained"  onClick={() => createGroup(props.segment.id)} color="primary">Create new group</Button>
            </AccordionDetails>
        </Accordion>
    );
}