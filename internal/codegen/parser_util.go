package main

func (doc KabusAPIDocument) TypedArrays() []SchemaDef {
	typedArrays := make([]SchemaDef, 0, len(doc.Schemas))
	for _, d := range doc.Schemas {
		if _, ok := d.Type.(ArrayDef); ok {
			typedArrays = append(typedArrays, d)
		}
	}
	return typedArrays
}

func (doc KabusAPIDocument) TypedStructs() []StructDef {
	typedStructs := make([]StructDef, 0, len(doc.Schemas)*2)
	for _, d := range doc.Schemas {
		if a, ok := d.Type.(ArrayDef); ok {
			if ss, ok := a.Elem.(StructDef); ok {
				typedStructs = append(typedStructs, ss)
				for _, p := range ss.Properties {
					if ss, ok := p.Type.(StructDef); ok {
						typedStructs = append(typedStructs, ss)
					}
					if a, ok := p.Type.(ArrayDef); ok {
						if ss, ok := a.Elem.(StructDef); ok {
							typedStructs = append(typedStructs, ss)
						}
					}
				}
				continue
			}
			continue
		}
		if s, ok := d.Type.(StructDef); ok {
			typedStructs = append(typedStructs, s)
			for _, p := range s.Properties {
				if ss, ok := p.Type.(StructDef); ok {
					typedStructs = append(typedStructs, ss)
				}
				if a, ok := p.Type.(ArrayDef); ok {
					if ss, ok := a.Elem.(StructDef); ok {
						typedStructs = append(typedStructs, ss)
					}
				}
			}
		}
	}
	return typedStructs
}
