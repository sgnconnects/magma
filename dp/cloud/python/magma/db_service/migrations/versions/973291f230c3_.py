"""empty message

Revision ID: 973291f230c3
Revises: 0c0edb6fd89b
Create Date: 2022-04-04 16:16:50.441833

"""
import sqlalchemy as sa
from alembic import op
from sqlalchemy.dialects import postgresql

# revision identifiers, used by Alembic.
revision = '973291f230c3'
down_revision = '0c0edb6fd89b'
branch_labels = None
depends_on = None


def upgrade():
    """
    Run upgrade
    """
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_constraint('requests_state_id_fkey', 'requests', type_='foreignkey')
    op.drop_column('requests', 'state_id')
    op.drop_table('responses')
    op.drop_table('request_states')
    # ### end Alembic commands ###


def downgrade():
    """
    Run downgrade
    """
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('requests', sa.Column('state_id', sa.INTEGER(), autoincrement=False, nullable=True))
    op.create_table(
        'request_states',
        sa.Column('id', sa.INTEGER(), server_default=sa.text("nextval('request_states_id_seq'::regclass)"), autoincrement=True, nullable=False),
        sa.Column('name', sa.VARCHAR(), autoincrement=False, nullable=False),
        sa.PrimaryKeyConstraint('id', name='request_states_pkey'),
        sa.UniqueConstraint('name', name='request_states_name_key'),
        postgresql_ignore_search_path=False,
    )
    op.create_foreign_key('requests_state_id_fkey', 'requests', 'request_states', ['state_id'], ['id'], ondelete='CASCADE')
    op.create_table(
        'responses',
        sa.Column('id', sa.INTEGER(), autoincrement=True, nullable=False),
        sa.Column('request_id', sa.INTEGER(), autoincrement=False, nullable=True),
        sa.Column('grant_id', sa.INTEGER(), autoincrement=False, nullable=True),
        sa.Column('response_code', sa.INTEGER(), autoincrement=False, nullable=False),
        sa.Column('created_date', postgresql.TIMESTAMP(timezone=True), server_default=sa.text('statement_timestamp()'), autoincrement=False, nullable=False),
        sa.Column('payload', postgresql.JSON(astext_type=sa.Text()), autoincrement=False, nullable=True),
        sa.ForeignKeyConstraint(['grant_id'], ['grants.id'], name='responses_grant_id_fkey', ondelete='CASCADE'),
        sa.ForeignKeyConstraint(['request_id'], ['requests.id'], name='responses_request_id_fkey', ondelete='CASCADE'),
        sa.PrimaryKeyConstraint('id', name='responses_pkey'),
    )
    # ### end Alembic commands ###
