# Cloud Native Indonesia Meetups

Repository to gather all meetup information and slides from Kubernetes and Cloud Native Indonesia meetups:

* [Bandung](bandung/README.md) ([meetup group](https://www.meetup.com/Kubernetes-and-Cloud-Native-Bandung))
  * Aditya Rachman Putra [@banditelol](https://github.com/banditelol), [Halofina](https://www.halofina.id)
  * Giri Kuncoro [@girikuncoro](https://github.com/girikuncoro), [Gojek](https://gojek.io)
  * Iqbal Syamil [@2pai](https://github.com/2pai), [Telkom University](https://telkomuniversity.ac.id)
  * Joshua Bezaleel [@joshuabezaleel](https://github.com/joshuabezaleel), [Institut Teknologi Bandung](https://www.itb.ac.id)
  * Zufar Dhiyaulhaq [@zufardhiyaulhaq](https://github.com/zufardhiyaulhaq), [Telkom University](https://telkomuniversity.ac.id)

* [Jakarta/Kubernetes](jakarta/kubernetes/README.md) ([meetup group](https://www.meetup.com/jakarta-kubernetes))
  * Eufrat Tsaqib [@eufat](https://github.com/eufat), [Nodeflux](https://www.nodeflux.io)
  * Giri Kuncoro [@girikuncoro](https://github.com/girikuncoro), [Gojek](https://gojek.io)
  * Imre Nagi [@imrenagi](https://github.com/imrenagi), [Google Developer Expert](https://developers.google.com/community/experts)
  * Iqbal Farabi [@qbl](https://github.com/qbl), [Gojek](https://gojek.io)
  * Irvi Aini [@irvifa](https://github.com/irvifa), [Traveloka](https://www.traveloka.com)

* [Jakarta/Cloud Native](jakarta/cloud-native/README.md) ([meetup group](https://www.meetup.com/Microservice-JKT))
  * Armand Caesario [@mandocaesar](http://github.com/mandocaesar), [Kata.ai](https://kata.ai)
  * Imre Nagi [@imrenagi](https://github.com/imrenagi), [Google Developer Expert](https://developers.google.com/community/experts)
  * Prakash Divyy [@prakashdivyy](http://github.com/prakashdivyy), [Kata.ai](https://kata.ai)

* [Yogyakarta](yogyakarta/README.md) ([meetup group](https://www.meetup.com/Kubernetes-and-Cloud-Native-Yogyakarta))
  * Agastyo Satriaji Idam [@satriajidam](https://github.com/satriajidam), [Ruangguru](http://www.ruangguru.com/)
  * Ary Dwi Marta P [@arydwimarta](https://github.com/arydwimarta), [Mamikos](https://mamikos.com/)
  * Deny Prasetyo [@jasoet](http://github.com/jasoet), [Gojek](https://gojek.io)
  * Manggala Pramuditya Wiryawan [@wiryawan46](https://github.com/wiryawan46), [Qiscus](https://www.qiscus.com/id)
  * Wahyuni Puji [@wahyuni-pj](https://github.com/wahyuni-pj), [Hilotech](http://hilotech.co.id/)

## Join our Community!

### Telegram

To facilitate and help each other in between meetups and different geographical locations, we have set up joined Telegram Groups.
These groups are also used for technical discussion around Kubernetes and Cloud Native topics.

In order to join, go to [@kubernetesindonesia](https://t.me/kubernetesindonesia) and [@cloudnativeid](https://t.me/microserviceid).
You can ask anything related to [CNCF open source projects](http://l.cncf.io/), and help each other's issues.

### Speaking Opportunities

If you'd like to speak at a meetup, please join our telegram group and mention **@girikuncoro** or **@imrenagi**, or fill [this form](https://goo.gl/forms/8UU0UgExUCqDMdp62).

## How to maintain this repo

Update `meetup.yaml` in your city with latest meetup info
```sh
cd jakarta/kubernetes
cat <<EOF >> meetup.yaml
  "20180110":
    title: "#1: Intro to Kubernetes and Knative"
    presentations:
    - duration: 30m0s
      recording: https://www.youtube.com/watch?v=DZQOgIWN1pE
      slides: https://bit.ly/knative
      speakers:
      - pahleviauliya
      title: Exploring KNative
EOF
```

Generate the markdown file
```sh
make
```
