// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/migrate"

	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/merchant"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/payment"
	"gitlab.com/mcuc/monorepo/backend/natasha/pkg/ent/voucher"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Merchant is the client for interacting with the Merchant builders.
	Merchant *MerchantClient
	// Payment is the client for interacting with the Payment builders.
	Payment *PaymentClient
	// Voucher is the client for interacting with the Voucher builders.
	Voucher *VoucherClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Merchant = NewMerchantClient(c.config)
	c.Payment = NewPaymentClient(c.config)
	c.Voucher = NewVoucherClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Merchant: NewMerchantClient(cfg),
		Payment:  NewPaymentClient(cfg),
		Voucher:  NewVoucherClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:   cfg,
		Merchant: NewMerchantClient(cfg),
		Payment:  NewPaymentClient(cfg),
		Voucher:  NewVoucherClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Merchant.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Merchant.Use(hooks...)
	c.Payment.Use(hooks...)
	c.Voucher.Use(hooks...)
}

// MerchantClient is a client for the Merchant schema.
type MerchantClient struct {
	config
}

// NewMerchantClient returns a client for the Merchant from the given config.
func NewMerchantClient(c config) *MerchantClient {
	return &MerchantClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `merchant.Hooks(f(g(h())))`.
func (c *MerchantClient) Use(hooks ...Hook) {
	c.hooks.Merchant = append(c.hooks.Merchant, hooks...)
}

// Create returns a create builder for Merchant.
func (c *MerchantClient) Create() *MerchantCreate {
	mutation := newMerchantMutation(c.config, OpCreate)
	return &MerchantCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Merchant entities.
func (c *MerchantClient) CreateBulk(builders ...*MerchantCreate) *MerchantCreateBulk {
	return &MerchantCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Merchant.
func (c *MerchantClient) Update() *MerchantUpdate {
	mutation := newMerchantMutation(c.config, OpUpdate)
	return &MerchantUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MerchantClient) UpdateOne(m *Merchant) *MerchantUpdateOne {
	mutation := newMerchantMutation(c.config, OpUpdateOne, withMerchant(m))
	return &MerchantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MerchantClient) UpdateOneID(id int64) *MerchantUpdateOne {
	mutation := newMerchantMutation(c.config, OpUpdateOne, withMerchantID(id))
	return &MerchantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Merchant.
func (c *MerchantClient) Delete() *MerchantDelete {
	mutation := newMerchantMutation(c.config, OpDelete)
	return &MerchantDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *MerchantClient) DeleteOne(m *Merchant) *MerchantDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *MerchantClient) DeleteOneID(id int64) *MerchantDeleteOne {
	builder := c.Delete().Where(merchant.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MerchantDeleteOne{builder}
}

// Query returns a query builder for Merchant.
func (c *MerchantClient) Query() *MerchantQuery {
	return &MerchantQuery{config: c.config}
}

// Get returns a Merchant entity by its id.
func (c *MerchantClient) Get(ctx context.Context, id int64) (*Merchant, error) {
	return c.Query().Where(merchant.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MerchantClient) GetX(ctx context.Context, id int64) *Merchant {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MerchantClient) Hooks() []Hook {
	return c.hooks.Merchant
}

// PaymentClient is a client for the Payment schema.
type PaymentClient struct {
	config
}

// NewPaymentClient returns a client for the Payment from the given config.
func NewPaymentClient(c config) *PaymentClient {
	return &PaymentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `payment.Hooks(f(g(h())))`.
func (c *PaymentClient) Use(hooks ...Hook) {
	c.hooks.Payment = append(c.hooks.Payment, hooks...)
}

// Create returns a create builder for Payment.
func (c *PaymentClient) Create() *PaymentCreate {
	mutation := newPaymentMutation(c.config, OpCreate)
	return &PaymentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Payment entities.
func (c *PaymentClient) CreateBulk(builders ...*PaymentCreate) *PaymentCreateBulk {
	return &PaymentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Payment.
func (c *PaymentClient) Update() *PaymentUpdate {
	mutation := newPaymentMutation(c.config, OpUpdate)
	return &PaymentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PaymentClient) UpdateOne(pa *Payment) *PaymentUpdateOne {
	mutation := newPaymentMutation(c.config, OpUpdateOne, withPayment(pa))
	return &PaymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PaymentClient) UpdateOneID(id int64) *PaymentUpdateOne {
	mutation := newPaymentMutation(c.config, OpUpdateOne, withPaymentID(id))
	return &PaymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Payment.
func (c *PaymentClient) Delete() *PaymentDelete {
	mutation := newPaymentMutation(c.config, OpDelete)
	return &PaymentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PaymentClient) DeleteOne(pa *Payment) *PaymentDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PaymentClient) DeleteOneID(id int64) *PaymentDeleteOne {
	builder := c.Delete().Where(payment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PaymentDeleteOne{builder}
}

// Query returns a query builder for Payment.
func (c *PaymentClient) Query() *PaymentQuery {
	return &PaymentQuery{config: c.config}
}

// Get returns a Payment entity by its id.
func (c *PaymentClient) Get(ctx context.Context, id int64) (*Payment, error) {
	return c.Query().Where(payment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PaymentClient) GetX(ctx context.Context, id int64) *Payment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PaymentClient) Hooks() []Hook {
	return c.hooks.Payment
}

// VoucherClient is a client for the Voucher schema.
type VoucherClient struct {
	config
}

// NewVoucherClient returns a client for the Voucher from the given config.
func NewVoucherClient(c config) *VoucherClient {
	return &VoucherClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `voucher.Hooks(f(g(h())))`.
func (c *VoucherClient) Use(hooks ...Hook) {
	c.hooks.Voucher = append(c.hooks.Voucher, hooks...)
}

// Create returns a create builder for Voucher.
func (c *VoucherClient) Create() *VoucherCreate {
	mutation := newVoucherMutation(c.config, OpCreate)
	return &VoucherCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Voucher entities.
func (c *VoucherClient) CreateBulk(builders ...*VoucherCreate) *VoucherCreateBulk {
	return &VoucherCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Voucher.
func (c *VoucherClient) Update() *VoucherUpdate {
	mutation := newVoucherMutation(c.config, OpUpdate)
	return &VoucherUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VoucherClient) UpdateOne(v *Voucher) *VoucherUpdateOne {
	mutation := newVoucherMutation(c.config, OpUpdateOne, withVoucher(v))
	return &VoucherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VoucherClient) UpdateOneID(id int64) *VoucherUpdateOne {
	mutation := newVoucherMutation(c.config, OpUpdateOne, withVoucherID(id))
	return &VoucherUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Voucher.
func (c *VoucherClient) Delete() *VoucherDelete {
	mutation := newVoucherMutation(c.config, OpDelete)
	return &VoucherDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VoucherClient) DeleteOne(v *Voucher) *VoucherDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VoucherClient) DeleteOneID(id int64) *VoucherDeleteOne {
	builder := c.Delete().Where(voucher.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VoucherDeleteOne{builder}
}

// Query returns a query builder for Voucher.
func (c *VoucherClient) Query() *VoucherQuery {
	return &VoucherQuery{config: c.config}
}

// Get returns a Voucher entity by its id.
func (c *VoucherClient) Get(ctx context.Context, id int64) (*Voucher, error) {
	return c.Query().Where(voucher.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VoucherClient) GetX(ctx context.Context, id int64) *Voucher {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VoucherClient) Hooks() []Hook {
	return c.hooks.Voucher
}
