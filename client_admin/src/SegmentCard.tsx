import React from 'react';
import { Accordion, AccordionDetails, AccordionSummary, Typography, IconButton, Box } from "@material-ui/core";
import ExpandMoreIcon from '@material-ui/icons/ExpandMore';
import AddCircleIcon from '@material-ui/icons/AddCircle';
import { ParticipantDto, SegmentDto } from './model';
import { createGroup, useAsyncError } from './service';
import Participants from './Participants';

interface SegmentCardProps {
    onRefresh: Function;
    segment: SegmentDto;
    participants?: ParticipantDto[]
}
export default function SegmentCard(props: SegmentCardProps) {
    const throwError = useAsyncError();
    
    return(
        <Accordion style={{margin: '5px 0px'}}>
            <AccordionSummary
                expandIcon={<ExpandMoreIcon />}
                aria-controls="panel1a-content"
                id="panel1a-header">
                <Box display="flex" alignItems="center">
                    <Typography variant="h5" component="h2">{props.segment.name}</Typography>
                    <IconButton size="medium" color="primary" onClick={() => {
                        createGroup(props.segment.id)
                            .then(() => props.onRefresh())
                            .catch((err) => throwError(err));
                    }}>
                        <AddCircleIcon />
                    </IconButton>
                </Box>
                
            </AccordionSummary>
            <AccordionDetails>
                <Participants participants = {props.participants}></Participants>
            </AccordionDetails>
        </Accordion>
    );
}