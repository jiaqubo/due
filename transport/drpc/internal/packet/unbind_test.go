package packet_test

import (
	"github.com/dobyte/due/v2/transport/drpc/internal/packet"
	"testing"
)

func TestUnbindPacker_Req(t *testing.T) {
	p := packet.NewUnbindPacker()

	buf, err := p.PackReq(1, 2)
	if err != nil {
		t.Fatal(err)
	}

	seq, uid, err := p.UnpackReq(buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("seq: %v uid: %v", seq, uid)
}

func TestUnbindPacker_Res(t *testing.T) {
	p := packet.NewUnbindPacker()

	buf, err := p.PackRes(1, 2)
	if err != nil {
		t.Fatal(err)
	}

	code, err := p.UnpackRes(buf.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("code: %v", code)
}