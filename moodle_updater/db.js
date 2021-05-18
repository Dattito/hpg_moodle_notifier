const { Sequelize, DataTypes } = require('sequelize');

const sequelize = new Sequelize(
    process.env.POSTGRES_DATABASE, 
    process.env.POSTGRES_USER,
    process.env.POSTGRES_PASSWORD,
    {
        host: process.env.POSTGRES_HOST,
        dialect: 'postgres',
        logging: false
    }
); 


const Assignment = sequelize.define('assignments', {
    id: {
        type: DataTypes.UUID,
        primaryKey: true,
        allowNull: false,
        defaultValue: Sequelize.UUIDV4,

    },
    moodleToken: {
        type: DataTypes.STRING,
        allowNull: false
    },
    assignments: {
        type: DataTypes.JSON,
        defaultValue: []
    },
    phoneNumber: {
        type: DataTypes.STRING,
        allowNull: false
    }
}, {
    underscored: true
});


module.exports = {
    sequelize: sequelize,
    Assignment: Assignment
}