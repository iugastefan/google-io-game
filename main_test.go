package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_play(t *testing.T) {
	type args struct {
		input *ArenaUpdate
	}
	js := "{\"_links\":{\"self\":{\"href\":\"https://go-2valpakcma-uc.a.run.app\"}},\"arena\":{\"dims\":[15,11],\"state\":{\"https://cloudbowl-samples-go-3pmsmtgcxa-et.a.run.app\":{\"x\":0,\"y\":2,\"direction\":\"W\",\"wasHit\":false,\"score\":-421},\"https://cloudbowl-samples-go-bew3uhmmqa-uc.a.run.app\":{\"x\":0,\"y\":3,\"direction\":\"N\",\"wasHit\":false,\"score\":-63},\"https://cloudbowl-samples-go-p2w5bb6fmq-uc.a.run.app\":{\"x\":4,\"y\":7,\"direction\":\"N\",\"wasHit\":false,\"score\":-208},\"https://cloudbowl-samples-go-v6oclnpx4a-uc.a.run.app\":{\"x\":6,\"y\":5,\"direction\":\"E\",\"wasHit\":false,\"score\":-20},\"https://cloudbowl-samples-java-quarkus-jshbj5hpnq-uc.a.run.app/\":{\"x\":11,\"y\":3,\"direction\":\"W\",\"wasHit\":false,\"score\":-128},\"https://cloudbowl-samples-java-springboot-5izzvnkrda-uc.a.run.app\":{\"x\":7,\"y\":1,\"direction\":\"E\",\"wasHit\":false,\"score\":-588},\"https://cloudbowl-samples-java-springboot-c324mhlcwq-uc.a.run.app\":{\"x\":3,\"y\":0,\"direction\":\"E\",\"wasHit\":false,\"score\":-14},\"https://cloudbowl-samples-java-springboot-yok2adf4vq-uc.a.run.app\":{\"x\":1,\"y\":4,\"direction\":\"W\",\"wasHit\":false,\"score\":-72},\"https://cloudbowl-samples-kotlin-micronaut-2oj2jzknaq-uc.a.run.app\":{\"x\":14,\"y\":0,\"direction\":\"W\",\"wasHit\":false,\"score\":-402},\"https://cloudbowl-samples-kotlin-micronaut-gih7xfneea-uc.a.run.app\":{\"x\":6,\"y\":3,\"direction\":\"E\",\"wasHit\":false,\"score\":-237},\"https://cloudbowl-samples-kotlin-micronaut-pu5yz7r3za-uc.a.run.app\":{\"x\":13,\"y\":3,\"direction\":\"E\",\"wasHit\":true,\"score\":-110},\"https://cloudbowl-samples-kotlin-springboot-7jbzybi3ka-uc.a.run.app\":{\"x\":1,\"y\":3,\"direction\":\"S\",\"wasHit\":false,\"score\":-543},\"https://cloudbowl-samples-nodejs-bvuipgyz3a-uc.a.run.app\":{\"x\":4,\"y\":2,\"direction\":\"S\",\"wasHit\":false,\"score\":-191},\"https://cloudbowl-samples-nodejs-gsb2dgomia-uc.a.run.app\":{\"x\":1,\"y\":7,\"direction\":\"S\",\"wasHit\":false,\"score\":-162},\"https://cloudbowl-samples-nodejs-iizt7cteoq-uc.a.run.app\":{\"x\":4,\"y\":9,\"direction\":\"E\",\"wasHit\":false,\"score\":-335},\"https://cloudbowl-samples-nodejs-khu47mcuxq-uc.a.run.app\":{\"x\":6,\"y\":9,\"direction\":\"W\",\"wasHit\":false,\"score\":-173},\"https://cloudbowl-samples-nodejs-s7chn4fvnq-uc.a.run.app/\":{\"x\":13,\"y\":8,\"direction\":\"W\",\"wasHit\":true,\"score\":-7232},\"https://cloudbowl-samples-nodejs-spfct6ypia-uc.a.run.app\":{\"x\":12,\"y\":7,\"direction\":\"W\",\"wasHit\":false,\"score\":-288},\"https://cloudbowl-samples-python-eh5k564wga-uc.a.run.app\":{\"x\":7,\"y\":2,\"direction\":\"E\",\"wasHit\":false,\"score\":-352},\"https://cloudbowl-samples-python-i7bb3yosnq-uc.a.run.app\":{\"x\":12,\"y\":2,\"direction\":\"E\",\"wasHit\":true,\"score\":-125},\"https://cloudbowl-samples-python-ldxd6veezq-uc.a.run.app/\":{\"x\":14,\"y\":3,\"direction\":\"W\",\"wasHit\":false,\"score\":-396},\"https://cloudbowl-samples-python-qvmyutxtyq-ew.a.run.app/\":{\"x\":14,\"y\":10,\"direction\":\"S\",\"wasHit\":true,\"score\":-505},\"https://cloudbowl-samples-python-tp3mfao42a-uc.a.run.app\":{\"x\":0,\"y\":6,\"direction\":\"E\",\"wasHit\":false,\"score\":-86},\"https://go-2valpakcma-uc.a.run.app\":{\"x\":14,\"y\":8,\"direction\":\"S\",\"wasHit\":false,\"score\":5526},\"https://kotlin-springboot-m2godeacfq-uc.a.run.app\":{\"x\":12,\"y\":3,\"direction\":\"N\",\"wasHit\":false,\"score\":762},\"https://nodejs-gvbubvhyeq-uc.a.run.app\":{\"x\":10,\"y\":8,\"direction\":\"E\",\"wasHit\":false,\"score\":6363}}}}\""
	var au ArenaUpdate
	json.Unmarshal([]byte(js), &au)

	type test struct {
		name string
		args args
		want string
	}
	var tests []test
	for i := 0; i < 100000; i++ {
		t := test{fmt.Sprintf("t%d", i), args{&au}, "T"}
		tests = append(tests, t)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := play(tt.args.input); got != tt.want {
				t.Errorf("play() = %v, want %v", got, tt.want)
			}
		})
	}
}
