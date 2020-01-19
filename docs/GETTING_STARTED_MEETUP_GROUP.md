# Create a new Meetup Group

Kubernetes and Cloud Native Indonesia is an organisation for cloud native (CNCF) Meetup Groups in Indonesia to collaborate.
Participating Meetup Groups can apply for CNCF membership and become official Cloud Native Computing Foundation Meetup Groups.
The details for applying can be found at [https://github.com/cncf/meetups#how-to-apply](https://github.com/cncf/meetups#how-to-apply)
once the group is created and has proven to be an active group. The catch is you need to maintain payment subscription through credit card
for the meetup group until you have done 3 meetups. If you want to avoid this, reach out to CNCF Ambassador **@girikuncoro** on 
[Twitter](https://twitter.com/girikuncoro) or [Telegram](https://t.me/girikuncoro) to let him create one for you for **free**.

Below you can find a “how to” guide for getting started with the creation of a kubernetes and cloud native meetup group.
Please let [@girikuncoro](https://twitter.com/girikuncoro) know your interest, so that he can invite you to private meetup organizers group.

## Prepare the Meetup Group creation

The premise is that you want to create a Meetup Group in a city in Indonesia, where no meetups connected to Cloud Native Indonesia
already exist. The easiest way to see what meetups already exist under the Cloud Native Indonesia brand, is probably to look at the
[list](https://github.com/cloudnative-id/meetups/blob/master/README.md).

If you do not find your city in the list please suggest the creation of a new meetup in your city in our [Telegram group](https://t.me/kubernetesindonesia) (mention @girikuncoro or @imrenagi)
and there will be response to that suggestion and help you getting started.

Once the decision for the creation of the meetup group has been decided, there are a couple of things that needs to be addressed.

## Create the Meetup on Meetup.com

*(You can skip this step by asking @girikuncoro to create one for you, but feel free if you want to do it yourself)*

Go to [https://meetup.com](https://meetup.com) and create an account there. Once you have that you can continue with the remaining tasks.

* Create the meetup group on meetup.com
* The naming used is Kubernetes-Cloud-Native-`City`, because people know about Kubernetes, but not Cloud Native yet. Jakarta is exception because we started first.
* Find some local people that will help you with the creation of the meetup.com group
* Follow the guides for creating a new [meetup group](https://help.meetup.com/hc/en-us/articles/360002882111-Starting-a-Meetup-group).
* Follow the guide for creating the [first event](https://help.meetup.com/hc/en-us/articles/360002881251).
* Description of the group can be taken from the examples of [Jakarta Kubernetes](https://www.meetup.com/jakarta-kubernetes/), [Microservice JKT/Cloud Native User Group](https://www.meetup.com/Microservice-JKT/), or [Kubernetes and Cloud Native Bandung](https://www.meetup.com/Kubernetes-and-Cloud-Native-Bandung/).
* Remember to use the [artwork](https://github.com/cloudnative-id/artwork) for the Meetup group and the event.

---

The following content will be geared towards maintaining the group by conducting monthly events. 

## Monthly Events

* It is encouraged to keep the group to **stay active** by hosting event once a month. It can be talks, workshops, study groups, etc. The number/quantity of participants/attendees should not be an issue.
* A form is provided to allow speakers to fill their talk proposals or hosts to fill their details. The form can be found at the CNCF-ID Google Drive.
* Usually, the format of the monthly event runs for around around 2 hours consisted of 2 speakers. The agenda can be seen as follows.

6:30 pm - 7:10 pm | Registration & Check In
7:10 pm - 7:20 pm | Welcoming & Introduction
7:20 pm - 7:50 pm | Speaker #1
7:50 pm - 8:20 pm | Speaker #2
8:20 pm - 8:35 pm | Q&A Panel
8:35 pm - 9:00 pm | Hallway Track

### Speakers

* Look for speakers that are interested in giving talks/presentations. You can reach to the local startups/companies, universities, or other tech communities. This [repository](https://github.com/rizafahmi/awesome-speakers-id) containing list of Indonesian speakers can be a good starting point.
* The topics should be around [Cloud Native landscape](https://landscape.cncf.io). Consult this to other organizers if you are not sure about the topic.

### Hosts/Venues

* Reach out to the local tech companies/startups or perhaps universities located around your city. Ask them whether they are interested in becoming the host of the monthly event.
* If they have an adequate area of the office, they will usually offer to use the auditorium/hall of the office. If not, you can offer them recommendations of public spaces/co-working spaces in your local city and they will sponsor in the form of the venue payment. 
* In return of sponsoring the event, offer them several benefits such as:
    * 1 (one) slot of speaker from the institution.
    * Logo of the institution will be shown in the publications below.
    * Several minutes at the event for the institutions to promote themselves.

### Publication and Meetup.com event

* After making sure that there are 2 speakers interested in giving presentation and 1 institution to host or sponsor the event, you can proceed to creating the publication of the particular month's event.
* Raise an issue in our [github.com/cloudnative-id/meetups](github.com/cloudnative-id/meetups) with the issue template that can be found [here](https://github.com/cloudnative-id/meetups/blob/master/.github/ISSUE_TEMPLATE/create-meetup-event.md).
* Made a simple poster consisting of the information of the events (date, time, location, speakers photo, speakers name, speakers organisations, name of the event). Example can be seen [here](https://www.meetup.com/jakarta-kubernetes/events/259186080/) or [here](https://drive.google.com/file/d/1eH9ofLdQ-YSSnPHtwaaureT4LYdyPEi1/view?usp=sharing). You can also, and encouraged, to use our poster generator [here](https://github.com/cloudnative-id/artwork/tree/master/poster/generator) with the tutorial written there. 
* Create a new event at the Meetup.com group with the details recommended as below:
    * #(number of how many events has been held) "short title of talk 1" & "short title of talk 2".
    * FAQ of the event.     
    * Welcoming message.
    * Location details of the venue.
    * For every talk:
        * Title of the talk.
        * Name of the speaker + role and organization.
        * Short bio of the speaker.
        * Abstract of the talk.
    * Agenda of the event.
    * Community leaders/organizers info and social media account.
* Example of the past events can be found at the Jakarta and Bandung meetups.
* You can set the RSVP's slot of 30%/40% more than the venue's capacity because there will usually be a drop rate of the attendees.
* It is encouraged to publish the event info around 1 week before the D-Day.
* It is recommended to record the event and speakers' talks and then upload it to video hosting site. The detailed instruction on this can be found [here](https://github.com/cloudnative-id/meetups/blob/master/docs/RECORDING_MEETUP_EVENT.md). 

### Community Slide
* A slide containing information for the community is usually made for each of the routine meetups. It usually contains the recent updates on Kubernetes/Cloud Native technologies, other related events, and news from the community.

### Centralized Repository and Document
* All of the documents belong to your city meetup group can be placed inside the CNCF-ID Google Drive under the folder of `<your city name>`.
* Do not forget to upload the speaker's deck to the [cloudnative-id/meetups](github.com/cloudnative-id/meetups) under the folder of `<your city name>` so that it can be accessed by everyone.

## Contributing to this guide

You are very welcome to contribute to this guide in order to make it as easy as possible to create a new meetup group and spread the good news about the Kubernetes and Cloud Native technologies.

Thank you for reading.

## References

[CNCF Meetup Best Practices](https://github.com/cncf/meetups#meetup-best-practices)