package znet

import (
	"000web/009zinx/utils"
	"000web/009zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (p *DataPack) GetHeadLen() uint32 {
	return 8
}

func (p *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *DataPack) Unpack(data []byte) (ziface.IMessage, error) {
	reader := bytes.NewReader(data)
	msg := &Message{}
	err := binary.Read(reader, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}
	if utils.GlobalObject.MaxPackageSize > 0 &&
		msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("data len too large")
	}
	return msg, nil
}
