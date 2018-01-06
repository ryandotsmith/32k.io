# Post Incident Analysis

## Objective
This template is designed to help operators understand what happened and how to prevent it from happening again. It is not designed to be used during an incident. After resolving a serious incident, be sure to capture any ephemeral data, then take some time to get rest and then return to conduct the analysis.

## Template:

### Synopsis

* Quick story to convey the incident
* Who was impacted
* Who was involved in response

### Timeline

* Relevant snippets from human and machine transcripts
* References to metrics and charts

### Probable Cause

* Contributing factors

### Remediation
* Technical and design changes
* Team and process changes
* Incident resolution changes

## Human Communication Transcripts
It is paramount to understand how your team interacted during the incident. Prevention comes in basically 2 flavors: operator improvements and technical redesign. Addressing any improvements to the operations of the service requires an in-depth understanding of how the operators interacted during the incident.

Collect transcripts from your team's chat tool (e.g. campfire, hipchat, irc) and curate them for long-term storage. Store them in a raw form and then pull up key events into your PIA. Don't forget to link your PIA to the raw transcripts.

## Machine Transcripts

*WARNING: This step may be time sensitive depending on your log storage behaviour.*

Collect all log data for any system that you have access to. Don't discriminate, storage is cheap and it is always best to error on the side of too much data. Aggregate these logs into the same store as your human transcripts.

Export all system metrics into static image files. Most systems resolute data overtime and it is important that you have the highest resolution data available to make any conclusions.
