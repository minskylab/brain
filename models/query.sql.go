// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package models

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createConversation = `-- name: CreateConversation :one
INSERT INTO conversations (
  phone_number, jid, context, conversation_buffer, conversation_summary, user_name
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, created_at, updated_at, phone_number, jid, context, conversation_buffer, conversation_summary, user_name
`

type CreateConversationParams struct {
	PhoneNumber         string
	Jid                 sql.NullString
	Context             sql.NullString
	ConversationBuffer  sql.NullString
	ConversationSummary sql.NullString
	UserName            sql.NullString
}

func (q *Queries) CreateConversation(ctx context.Context, arg CreateConversationParams) (Conversation, error) {
	row := q.db.QueryRowContext(ctx, createConversation,
		arg.PhoneNumber,
		arg.Jid,
		arg.Context,
		arg.ConversationBuffer,
		arg.ConversationSummary,
		arg.UserName,
	)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PhoneNumber,
		&i.Jid,
		&i.Context,
		&i.ConversationBuffer,
		&i.ConversationSummary,
		&i.UserName,
	)
	return i, err
}

const deleteConversation = `-- name: DeleteConversation :exec
DELETE FROM conversations
WHERE id = $1
`

func (q *Queries) DeleteConversation(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteConversation, id)
	return err
}

const getConversation = `-- name: GetConversation :one
SELECT id, created_at, updated_at, phone_number, jid, context, conversation_buffer, conversation_summary, user_name FROM conversations
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetConversation(ctx context.Context, id uuid.UUID) (Conversation, error) {
	row := q.db.QueryRowContext(ctx, getConversation, id)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PhoneNumber,
		&i.Jid,
		&i.Context,
		&i.ConversationBuffer,
		&i.ConversationSummary,
		&i.UserName,
	)
	return i, err
}

const getConversationByJid = `-- name: GetConversationByJid :one
SELECT id, created_at, updated_at, phone_number, jid, context, conversation_buffer, conversation_summary, user_name FROM conversations
WHERE jid = $1 LIMIT 1
`

func (q *Queries) GetConversationByJid(ctx context.Context, jid sql.NullString) (Conversation, error) {
	row := q.db.QueryRowContext(ctx, getConversationByJid, jid)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PhoneNumber,
		&i.Jid,
		&i.Context,
		&i.ConversationBuffer,
		&i.ConversationSummary,
		&i.UserName,
	)
	return i, err
}

const getConversationByPhoneNumber = `-- name: GetConversationByPhoneNumber :one
SELECT id, created_at, updated_at, phone_number, jid, context, conversation_buffer, conversation_summary, user_name FROM conversations
WHERE phone_number = $1 LIMIT 1
`

func (q *Queries) GetConversationByPhoneNumber(ctx context.Context, phoneNumber string) (Conversation, error) {
	row := q.db.QueryRowContext(ctx, getConversationByPhoneNumber, phoneNumber)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PhoneNumber,
		&i.Jid,
		&i.Context,
		&i.ConversationBuffer,
		&i.ConversationSummary,
		&i.UserName,
	)
	return i, err
}

const listConversations = `-- name: ListConversations :many
SELECT id, created_at, updated_at, phone_number, jid, context, conversation_buffer, conversation_summary, user_name FROM conversations
ORDER BY updated_at
`

func (q *Queries) ListConversations(ctx context.Context) ([]Conversation, error) {
	rows, err := q.db.QueryContext(ctx, listConversations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Conversation
	for rows.Next() {
		var i Conversation
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.PhoneNumber,
			&i.Jid,
			&i.Context,
			&i.ConversationBuffer,
			&i.ConversationSummary,
			&i.UserName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateConversationBuffer = `-- name: UpdateConversationBuffer :one
UPDATE conversations
SET conversation_buffer = $1
WHERE id = $2
RETURNING id, created_at, updated_at, phone_number, jid, context, conversation_buffer, conversation_summary, user_name
`

type UpdateConversationBufferParams struct {
	ConversationBuffer sql.NullString
	ID                 uuid.UUID
}

func (q *Queries) UpdateConversationBuffer(ctx context.Context, arg UpdateConversationBufferParams) (Conversation, error) {
	row := q.db.QueryRowContext(ctx, updateConversationBuffer, arg.ConversationBuffer, arg.ID)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PhoneNumber,
		&i.Jid,
		&i.Context,
		&i.ConversationBuffer,
		&i.ConversationSummary,
		&i.UserName,
	)
	return i, err
}

const updateConversationContext = `-- name: UpdateConversationContext :one
UPDATE conversations
SET context = $1
WHERE id = $2
RETURNING id, created_at, updated_at, phone_number, jid, context, conversation_buffer, conversation_summary, user_name
`

type UpdateConversationContextParams struct {
	Context sql.NullString
	ID      uuid.UUID
}

func (q *Queries) UpdateConversationContext(ctx context.Context, arg UpdateConversationContextParams) (Conversation, error) {
	row := q.db.QueryRowContext(ctx, updateConversationContext, arg.Context, arg.ID)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PhoneNumber,
		&i.Jid,
		&i.Context,
		&i.ConversationBuffer,
		&i.ConversationSummary,
		&i.UserName,
	)
	return i, err
}

const updateConversationSummary = `-- name: UpdateConversationSummary :one
UPDATE conversations
SET conversation_summary = $1
WHERE id = $2
RETURNING id, created_at, updated_at, phone_number, jid, context, conversation_buffer, conversation_summary, user_name
`

type UpdateConversationSummaryParams struct {
	ConversationSummary sql.NullString
	ID                  uuid.UUID
}

func (q *Queries) UpdateConversationSummary(ctx context.Context, arg UpdateConversationSummaryParams) (Conversation, error) {
	row := q.db.QueryRowContext(ctx, updateConversationSummary, arg.ConversationSummary, arg.ID)
	var i Conversation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.PhoneNumber,
		&i.Jid,
		&i.Context,
		&i.ConversationBuffer,
		&i.ConversationSummary,
		&i.UserName,
	)
	return i, err
}
