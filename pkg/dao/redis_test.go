package dao

import (
	"testing"

	"dispatcher/pkg/config"
	"dispatcher/pkg/log"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const hashId = "qrRodnpQNbQ"

func TestEnv_SetRedis(t *testing.T) {
	config.SetConfig()
	logger, _ := log.InitLog()
	redisConn := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: "",
		DB:       0,
	})

	type fields struct {
		Logger *zap.Logger
		Redis  *redis.Client
	}
	type args struct {
		fileContent string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "SetFileContentToRedis",
			fields: fields{
				Logger: logger,
				Redis:  redisConn,
			},
			args: args{
				fileContent: "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",
			},
			want:    hashId,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Env{
				Logger: tt.fields.Logger,
				Redis:  tt.fields.Redis,
			}
			got, err := e.SetRedis(tt.args.fileContent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Env.SetRedis() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Env.SetRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}
