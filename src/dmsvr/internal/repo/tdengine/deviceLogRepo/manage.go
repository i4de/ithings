package deviceLogRepo

import (
	"context"
	"fmt"
)

func (d DeviceLogRepo) InitProduct(ctx context.Context, productID string) error {
	sql := fmt.Sprintf("CREATE STABLE IF NOT EXISTS %s "+
		"(`ts` timestamp,`content` BINARY(5000),`topic` BINARY(500), `action` BINARY(100),"+
		" `request_id` BINARY(100), `trance_id` BINARY(100), `result_type` BIGINT)"+
		"TAGS (device_name BINARY(50));",
		getLogStableName(productID))
	if _, err := d.t.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (d DeviceLogRepo) DropProduct(ctx context.Context, productID string) error {
	sql := fmt.Sprintf("drop stable if exists %s;", getLogStableName(productID))
	if _, err := d.t.Exec(sql); err != nil {
		return err
	}
	return nil
}

func (d DeviceLogRepo) DropDevice(ctx context.Context, productID string, deviceName string) error {
	sql := fmt.Sprintf("drop table if exists %s;", getLogTableName(productID, deviceName))
	if _, err := d.t.Exec(sql); err != nil {
		return err
	}
	return nil
}
