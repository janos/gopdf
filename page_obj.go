package gopdf

import (
	"bytes"
	//"fmt"
)

type PageObj struct { //impl IObj
	buffer          bytes.Buffer
	Contents        string
	ResourcesRelate string
}

func (p *PageObj) Init(funcGetRoot func() *GoPdf) {

}

func (p *PageObj) Build() error {

	p.buffer.WriteString("<<\n")
	p.buffer.WriteString("  /Type /" + p.GetType() + "\n")
	p.buffer.WriteString("  /Parent 2 0 R\n")
	p.buffer.WriteString("  /Resources " + p.ResourcesRelate + "\n")
	/*me.buffer.WriteString("    /Font <<\n")
	i := 0
	max := len(me.Realtes)
	for i < max {
		realte := me.Realtes[i]
		me.buffer.WriteString(fmt.Sprintf("      /F%d %d 0 R\n",realte.CountOfFont + 1, realte.IndexOfObj + 1))
		i++
	}
	me.buffer.WriteString("    >>\n")*/
	//me.buffer.WriteString("  >>\n")
	p.buffer.WriteString("  /Contents " + p.Contents + "\n") //sample  Contents 8 0 R
	p.buffer.WriteString(">>\n")
	return nil
}

func (p *PageObj) GetType() string {
	return "Page"
}

func (p *PageObj) GetObjBuff() *bytes.Buffer {
	return &(p.buffer)
}
