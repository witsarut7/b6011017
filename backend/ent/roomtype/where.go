// Code generated by entc, DO NOT EDIT.

package roomtype

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/witsarut7/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ROOMPRICE applies equality check predicate on the "ROOMPRICE" field. It's identical to ROOMPRICEEQ.
func ROOMPRICE(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldROOMPRICE), v))
	})
}

// ROOMPRICEEQ applies the EQ predicate on the "ROOMPRICE" field.
func ROOMPRICEEQ(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldROOMPRICE), v))
	})
}

// ROOMPRICENEQ applies the NEQ predicate on the "ROOMPRICE" field.
func ROOMPRICENEQ(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldROOMPRICE), v))
	})
}

// ROOMPRICEIn applies the In predicate on the "ROOMPRICE" field.
func ROOMPRICEIn(vs ...int) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldROOMPRICE), v...))
	})
}

// ROOMPRICENotIn applies the NotIn predicate on the "ROOMPRICE" field.
func ROOMPRICENotIn(vs ...int) predicate.Roomtype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Roomtype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldROOMPRICE), v...))
	})
}

// ROOMPRICEGT applies the GT predicate on the "ROOMPRICE" field.
func ROOMPRICEGT(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldROOMPRICE), v))
	})
}

// ROOMPRICEGTE applies the GTE predicate on the "ROOMPRICE" field.
func ROOMPRICEGTE(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldROOMPRICE), v))
	})
}

// ROOMPRICELT applies the LT predicate on the "ROOMPRICE" field.
func ROOMPRICELT(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldROOMPRICE), v))
	})
}

// ROOMPRICELTE applies the LTE predicate on the "ROOMPRICE" field.
func ROOMPRICELTE(v int) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldROOMPRICE), v))
	})
}

// HasPayment applies the HasEdge predicate on the "payment" edge.
func HasPayment() predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPaymentWith applies the HasEdge predicate on the "payment" edge with a given conditions (other predicates).
func HasPaymentWith(preds ...predicate.Payment) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PaymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PaymentTable, PaymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Roomtype) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Roomtype) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Roomtype) predicate.Roomtype {
	return predicate.Roomtype(func(s *sql.Selector) {
		p(s.Not())
	})
}