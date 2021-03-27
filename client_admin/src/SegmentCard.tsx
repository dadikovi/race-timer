import React from 'react';
import { Accordion, AccordionDetails, AccordionSummary, Typography, Button } from "@material-ui/core";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import { SegmentDto } from './model';
import { createGroup, useAsyncError } from './service';

interface SegmentCardProps {
    onRefresh: Function;
    segment: SegmentDto;
}
export default function SegmentCard(props: SegmentCardProps) {
    const throwError = useAsyncError();
    
    return(
        <Accordion style={{margin: '5px 0px'}}>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel1a-content"
                id="panel1a-header">
                <Typography>{props.segment.name}</Typography>
            </AccordionSummary>
            <AccordionDetails>
                <Button variant="contained" onClick={() => {
                    createGroup(props.segment.id)
                        .then(() => props.onRefresh())
                        .catch((err) => throwError(err));
                }} color="primary">Create new group</Button>
            </AccordionDetails>
        </Accordion>
    );
}